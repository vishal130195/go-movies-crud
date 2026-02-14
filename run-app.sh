#!/bin/bash

echo "🚀 Starting Movies CRUD Application with separate servers..."
echo ""

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go first."
    exit 1
fi

# Check if we're in the right directory
if [ ! -f "go.mod" ]; then
    echo "❌ Please run this script from the project root directory."
    exit 1
fi

echo "📦 Installing dependencies..."
go mod tidy

echo ""
echo "🔧 Starting Backend API Server (Port 8000)..."
# Start the Go backend server in the background
go run cmd/api/main.go &
BACKEND_PID=$!

# Wait a moment for the backend server to start
sleep 3

echo "✅ Backend server started at: http://localhost:8000"
echo ""

echo "🌐 Starting UI Server (Port 3000)..."
# Start the UI server in the background
cd ui
go run server.go &
UI_PID=$!
cd ..

# Wait a moment for the UI server to start
sleep 2

echo "✅ UI server started at: http://localhost:3000"
echo ""

echo "🎉 Application is ready!"
echo ""
echo "📝 Available Services:"
echo "  🔗 UI Application: http://localhost:3000"
echo "  🔗 Direct HTML: http://localhost:3000/index.html"
echo "  🔗 Backend API: http://localhost:8000"
echo ""

# Try to open the UI in the browser
echo "🌐 Opening application in browser..."
if command -v open &> /dev/null; then
    # macOS
    open "http://localhost:3000" 2>/dev/null
elif command -v xdg-open &> /dev/null; then
    # Linux
    xdg-open "http://localhost:3000"
elif command -v start &> /dev/null; then
    # Windows
    start "http://localhost:3000"
else
    echo "Please manually open http://localhost:3000 in your browser"
fi

echo ""
echo "🛑 To stop both servers, press Ctrl+C"
echo ""

# Function to cleanup both processes
cleanup() {
    echo ""
    echo "🛑 Stopping servers..."
    kill $BACKEND_PID 2>/dev/null
    kill $UI_PID 2>/dev/null
    echo "✅ Servers stopped."
    exit 0
}

# Wait for Ctrl+C to stop both servers
trap cleanup SIGINT

# Keep the script running
wait

