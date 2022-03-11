package cronTask

const (
	STAT_IS_CAN_EXCUTE int =iota+1
	STAT_IS_NOT_EXCUTE
	STAT_IS_READ_BUT_NOTIFY
	STAT_IS_DELETED
)

const JOB_KEY_PREFIX string = "delay_job_prefix:"
