package jobs

import (
	"errors"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/sirupsen/logrus"

	"go.mikenewswanger.com/kube-ci/kube-ci/jobs/notifiers"
	"go.mikenewswanger.com/utilities/filesystem"
)

// Load imports jobs from a datastore
func Load(datastore string) (map[string]*Job, map[string]*notifiers.Notification, error) {
	// split[0] is datastoreType; //split[1] is connection string
	split := strings.SplitN(datastore, ":", 2)
	if len(split) != 2 {
		logrus.Fatal("Invalid Datastore Format")
	}

	var jobs map[string]*Job
	var notifications map[string]*notifiers.Notification
	var err error
	switch split[0] {
	case "filesystem":
		jobs, notifications, err = loadFromFilesystem(split[1])
	default:
		panic("Datastore not implented: " + split[0])
	}
	if err == nil {
		// Jobs and notifiers loaded successfully; validate and map them
		for _, j := range jobs {
			for _, n := range j.Notifiers {
				for _, t := range n {
					err = t.Bind(notifications)
					if err != nil {
						logrus.Error(err)
						break
					}
				}
			}
		}
	}
	return jobs, notifications, err
}

// Load from the filesystem
func loadFromFilesystem(path string) (map[string]*Job, map[string]*notifiers.Notification, error) {
	if !filesystem.IsDirectory(path) {
		return nil, nil, errors.New("Specified path (\"" + path + "\") does not exist or is not accessible")
	}
	n, err := loadNotificationsFromFilesystem(path + "/notifiers")
	if err != nil {
		return nil, nil, err
	}
	j, err := loadJobsFromFilesystem(path + "/jobs")
	if err != nil {
		return nil, nil, err
	}
	return j, n, nil
}

func loadJobsFromFilesystem(path string) (map[string]*Job, error) {
	if !filesystem.IsDirectory(path) {
		return nil, errors.New("Specified path (\"" + path + "\") does not exist or is not accessible")
	}
	directoryContents, err := filesystem.GetDirectoryContents(path)
	if err != nil {
		return nil, err
	}
	j := map[string]*Job{}
	for _, f := range directoryContents {
		job := &Job{}
		f = path + "/" + f
		fc, err := filesystem.LoadFileBytes(f)
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(fc, job)
		if err != nil {
			return nil, err
		}
		j[job.Namespace+"."+job.Name] = job
	}

	return j, nil
}

func loadNotificationsFromFilesystem(path string) (map[string]*notifiers.Notification, error) {
	if !filesystem.IsDirectory(path) {
		return nil, errors.New("Specified path (\"" + path + "\") does not exist or is not accessible")
	}
	directoryContents, err := filesystem.GetDirectoryContents(path)
	if err != nil {
		return nil, err
	}
	var notification *notifiers.Notification
	n := map[string]*notifiers.Notification{}
	for _, f := range directoryContents {
		f = path + "/" + f
		fc, err := filesystem.LoadFileBytes(f)
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(fc, &notification)
		if err != nil {
			return nil, err
		}
		n[notification.Namespace+"."+notification.Name] = notification
	}

	return n, nil
}
