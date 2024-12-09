@echo off

:: Start the Go backend
echo Starting Go backend...
start cmd /k "cd C:\Users\stefa\Desktop\Vue\leaderboard-app\backend && go run main.go"

:: Start the Vue.js frontend
echo Starting Vue.js frontend...
start cmd /k "cd C:\Users\stefa\Desktop\Vue\leaderboard-app\frontend && npm run dev"

echo Both the backend and frontend have been started!
