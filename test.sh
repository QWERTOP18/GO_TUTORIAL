#!/bin/bash
options=($(ls -d ex*/ 2>/dev/null) "Exit")

echo "Please choose a test case:"
select option in "${options[@]}"; do
    case $option in
        "Exit")
            echo "Exiting..."
            break
            ;;
        "")
            echo "Invalid option. Please select a valid test case."
            ;;
        *)
            if [ -d "$option" ]; then
                echo "Selected: $option"
                cd "$option" || exit
                # go.mod が存在しない場合、初期化
                if [ ! -f "go.mod" ]; then
                    go mod init "$(basename "$option")"
                fi
                go run .
                cd - || exit
            else
                echo "Invalid option. Please select a valid test case."
            fi
            ;;
    esac
done
