package jobs

import (
	"errors"

	"github.com/ghodss/yaml"

	"go.mikenewswanger.com/kube-ci/kube-ci/jobs/notifiers"
	"go.mikenewswanger.com/utilities/filesystem"
)

// Load imports jobs from a datastore
func Load(datastore string, path string) (map[string]*Job, map[string]*notifiers.Notification, error) {
	switch datastore {
	case "filesystem":
		return loadFromFilesystem(path)
	}
	panic("Could not load jobs")
}

// Load from the filesystem
func loadFromFilesystem(path string) (map[string]*Job, map[string]*notifiers.Notification, error) {
	fs := filesystem.Filesystem{}
	if !fs.IsDirectory(path) {
		return nil, nil, errors.New("Specified path (" + path + ") does not exist or is not accessible")
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
	var fs = filesystem.Filesystem{}
	if !fs.IsDirectory(path) {
		return nil, errors.New("Specified path (" + path + ") does not exist or is not accessible")
	}
	directoryContents, err := fs.GetDirectoryContents(path)
	if err != nil {
		return nil, err
	}
	var job *Job
	j := map[string]*Job{}
	for _, f := range directoryContents {
		f = path + "/" + f
		fc, err := fs.LoadFileBytes(f)
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(fc, &job)
		if err != nil {
			return nil, err
		}
		j[job.Namespace+"."+job.Name] = job
	}

	return j, nil
}

func loadNotificationsFromFilesystem(path string) (map[string]*notifiers.Notification, error) {
	var fs = filesystem.Filesystem{}
	if !fs.IsDirectory(path) {
		return nil, errors.New("Specified path (" + path + ") does not exist or is not accessible")
	}
	directoryContents, err := fs.GetDirectoryContents(path)
	if err != nil {
		return nil, err
	}
	var notification *notifiers.Notification
	n := map[string]*notifiers.Notification{}
	for _, f := range directoryContents {
		f = path + "/" + f
		fc, err := fs.LoadFileBytes(f)
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
