package lock

import (
	"os"
	"syscall"

	"github.com/bborbe/log"
)

type Lock interface {
	Lock() error
	Unlock() error
}

type lock struct {
	lockName string
	file     *os.File
}

var logger = log.DefaultLogger

func NewLock(lockName string) *lock {
	l := new(lock)
	l.lockName = lockName
	return l
}

func (l *lock) Lock() error {
	logger.Debug("try lock")
	var err error
	l.file, _ = os.Open(l.lockName)
	if l.file == nil {
		l.file, err = os.Create(l.lockName)
		if err != nil {
			logger.Debug("create lock file failed")
			return err
		}
	}
	err = syscall.Flock(int(l.file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		logger.Debug("lock fail, already locked")
		return err
	}
	logger.Debug("locked")
	return nil
}

func (l *lock) Unlock() error {
	logger.Debug("try unlock")
	var err error
	err = syscall.Flock(int(l.file.Fd()), syscall.LOCK_UN)
	if err != nil {
		logger.Debug("unlock failed")
		return err
	}
	logger.Debug("unlocked")
	return l.file.Close()
}
