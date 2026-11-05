# CYSE6225_assignment2

This README will detail how to run this Rest API that accepts uploads of only .xlsx files catagorizing them by a ID.

There are multiple commands that this Rest API can as mentioned above the Rest API can upload files as a POST method, these files can only be xlsx files meaning something like a .txt or a .pdf are unable to be accepted. There are two commands for the GET method which will either list all of the files on hand or one specific file via the file id. The PUT method can replace a specified file ID with a selected file. Lastly the DELETE method can delete selected files.

There are a few pre-requisites required to be able to run Docker needs to be installed with the Docker Desktop app, all the project files I provided, and bash or another command line program.

First things first ensure Docker Desktop is open and running before continuing anything else. I personally prefer running the shellscripts through the CMD but it can be run directly through the bash command line. I will be detailing the methods of running the base API first before going through the test files. I work on a windows computer so this will be continuing with that in mind. I go to the folder housing all of my files and open a CMD, in the CMD use the command "C:\Program Files\Git\bin\bash.exe" run_api.sh. This will execute preset Docker commands both building the image and running the html server in a container. The html server will automatically be hosted in: http://localhost:8000 with the message 404 Server not found being on the webpage.

Now in a separate CMD still running from the initial folder we can all the commands. I provided two file test_file.xlsx and test_file_2.xlsx to show off successful POST,GET,PUT, & DELETE functions. There is nothing stored in terms of data so I will begin with the POST method: in the new CMD use the input curl -i -F "file=@test_file.xlsx" http://localhost:8000/upload. This will upload the test_file.xlsx provided giving it the ID of 1. In the section of the command "file=@test_file.xlsx" replace everything after the "file=@" in order to upload either the test_file_2.xlsx or any other file imported.

For the sake of the tutorial use the command: curl -i -F "file=@test_file_2.xlsx" to upload another file. In order to see these files you need to invoke the GET method. In the CMD use curl -i http://localhost:8000/files to see every file uploaded or use curl -i http://localhost:8000/files/1 to see any file with the file ID of 1. 

the first command output should be show: 
File 1: test_file.xlsx
File 2: test_file_2.xlsx

the second command output should be show:
File 1: test_file.xlsx

Let's say you wanted to replace one of the files using the PUT method. For the sake of the example you wanted to replace test_file_2 with test_file (we will ignore the fact it becomes a duplicate of test_file for now). In the CMD use curl -i -X PUT -F "file=@test_file.xlsx" http://localhost:8000/files/2. The /2 after files will pinpoint the exact file to be replaced. Now if you run curl -i http://localhost:8000/files it should display:
File 1: test_file.xlsx
File 2: test_file.xlsx

Now for the DELETE method. Let's say you didn't want to have two test_file.xlsx's and you just wanted to keep the newly updated one. The command curl -i -X DELETE http://localhost:8000/files/1 will delete the File 1: test_file.xlsx leaving only the File 2: test_file.xlsx.  curl -i http://localhost:8000/files will confirm the file was deleted. 

Now to how to run the test scripts. Much like ensure Docker desktop is open before opening a CMD spawned from the folder holding all the project files run the shellscript using: "C:\Program Files\Git\bin\bash.exe" run_tests.sh. This will automatically build the Docker image and make a container. From here the tests will run automatically telling you if all the functions either pass or fail. There is an automatic workflow in GitHub that can be ran automatically when changes are pushed to Main or if you go to the Actions tab and press on the workflow Cl to find the Run workflow button. This will additionally send notifications to a personalized Slack channel on whether the build passes or fails. 


I wanted to have some works cited at the end of the README:

https://oneuptime.com/blog/post/2026-01-30-how-to-build-file-upload-service-in-go/view#mime-type-validation
https://roadmap.sh/golang/rest-api#step-3-implement-crud-operations-for-the-bookkeeping-application
https://reintech.io/blog/go-net-http-improvements-go-1-22-enhanced-routing-patterns
https://dev.to/neelp03/building-a-file-upload-service-in-go-34fj
https://www.gofaq.org/en/how-to-handle-file-uploads-in-a-go-http-server/
https://www.geeksforgeeks.org/go-language/golang-maps/
https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status
https://gobyexample.com/mutexes
https://speedscale.com/blog/testing-golang-with-httptest/
https://www.compilenrun.com/docs/language/go/go-testing/go-http-testing/
https://codesignal.com/learn/courses/automating-api-tests-with-go/lessons/testing-crud-operations-with-setup-and-teardown-in-go
I was having some issue with making the test files work so I did end up using Claude for the uploadRequest func in my testing files in addition to asking questions about how/why certain pieces of code work the way they are where I was introduced to concepts like mutexes which I've never used before.
