#!/bin/bash
echo "Please choose a test case:"
select option in ex** "Exit";  do
    case $option in
        "1")
            #番号で選べるようにする
            cd ex00
            #if *.mod doesnt exist
            go mod init ex00
            go run .
            break
            cd -
            ;;

        "Exit")
            echo "Exiting..."
            break
            ;;
        *)
            echo "Invalid option. Please select a valid test case."
            ;;
    esac
done
