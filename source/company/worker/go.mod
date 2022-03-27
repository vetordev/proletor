module bot

go 1.17

require job v0.0.0-00010101000000-000000000000

require github.com/robfig/cron/v3 v3.0.1 // indirect

replace job => ./../job
