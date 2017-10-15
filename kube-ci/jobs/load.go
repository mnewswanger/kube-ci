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
func Load(datastore string) (jobs map[string]*Job, notifications map[string]*notifiers.Notification, err error) {
	// split[0] is datastoreType; //split[1] is connection string
	split := strings.SplitN(datastore, ":", 2)
	if len(split) != 2 {
		logrus.Fatal("Invalid Datastore Format")
	}

	switch split[0] {
	case "filesystem":
		jobs, notifications, err = loadFromFilesystem(split[1])
	default:
		err = errors.New("Configuration Datastore is not supported")
		return
	}
	if err != nil {
		return
	}
	// Register notifications
	for _, n := range notifications {
		err = n.Register()
		if err != nil {
			return
		}
	}
	// Jobs and notifiers loaded successfully; validate and map them
	// Event fire type doesn't matter at this stage, so just focus on the notifications themselves
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
	return
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

func loadJobsFromFilesystem(path string) (jobs map[string]*Job, err error) {
	if !filesystem.IsDirectory(path) {
		return nil, errors.New("Specified path (\"" + path + "\") does not exist or is not accessible")
	}
	directoryContents, err := filesystem.GetDirectoryContents(path)
	if err != nil {
		return nil, err
	}
	jobs = map[string]*Job{}
	for _, f := range directoryContents {
		if strings.HasPrefix(f, ".") || f == "readme.md" {
			continue
		}
		f = path + "/" + f
		if filesystem.IsDirectory(f) {
			j, err := loadJobsFromFilesystem(f)
			if err != nil {
				return nil, err
			}
			for k, v := range j {
				jobs[k] = v
			}
			continue
		}
		fc, err := filesystem.LoadFileBytes(f)
		if err != nil {
			return nil, err
		}
		job := &Job{}
		err = yaml.Unmarshal(fc, job)
		if err != nil {
			return nil, err
		}
		jobs[job.Namespace+"."+job.Name] = job
	}

	return
}

func loadNotificationsFromFilesystem(path string) (notifications map[string]*notifiers.Notification, err error) {
	if !filesystem.IsDirectory(path) {
		return nil, errors.New("Specified path (\"" + path + "\") does not exist or is not accessible")
	}
	directoryContents, err := filesystem.GetDirectoryContents(path)
	if err != nil {
		return nil, err
	}
	notifications = map[string]*notifiers.Notification{}
	for _, f := range directoryContents {
		if strings.HasPrefix(f, ".") || f == "readme.md" {
			continue
		}
		f = path + "/" + f
		if filesystem.IsDirectory(f) {
			n, err := loadNotificationsFromFilesystem(f)
			if err != nil {
				return nil, err
			}
			for k, v := range n {
				notifications[k] = v
			}
			continue
		}
		fc, err := filesystem.LoadFileBytes(f)
		if err != nil {
			return nil, err
		}
		notification := &notifiers.Notification{}
		err = yaml.Unmarshal(fc, &notification)
		if err != nil {
			return nil, err
		}
		notifications[notification.Namespace+"."+notification.Name] = notification
	}

	return
}
