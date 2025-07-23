#!/bin/bash

# --- Detect the script's directory ---
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# --- Detect terminal ---
detect_terminal() {
  if [[ "$TERM_PROGRAM" == "iTerm.app" ]]; then
    echo "iTerm2"
  elif [[ "$TERM_PROGRAM" == "Apple_Terminal" ]]; then
    echo "Terminal"
  else
    parent=$(ps -o comm= -p $(ps -o ppid= -p $$))
    case "$parent" in
      gnome-terminal-*)
        echo "GNOME Terminal"
        ;;
      konsole)
        echo "Konsole"
        ;;
      *)
        echo "Unknown"
        ;;
    esac
  fi
}

terminal=$(detect_terminal)
echo "Detected terminal: $terminal"

# --- Run terminal-specific tab opening ---
if [[ "$terminal" == "iTerm2" ]]; then
  osascript <<EOF
tell application "iTerm2"
  create window with default profile
  tell current window
    tell current session
      write text "cd '$SCRIPT_DIR' && cd server/ && go run main.go"
      delay 1
    end tell
    create tab with default profile
    tell current session
      write text "cd '$SCRIPT_DIR' && cd web/ && npm run start"
    end tell
  end tell
end tell
EOF

elif [[ "$terminal" == "Terminal" ]]; then
  osascript <<EOF
tell application "Terminal"
  do script "cd '$SCRIPT_DIR' && cd server/ && go run main.go"
  delay 1
  do script "cd '$SCRIPT_DIR' && cd web && npm run start" in front window
end tell
EOF

else
  echo "Unsupported terminal emulator for tab automation. Please run the commands manually."
fi
