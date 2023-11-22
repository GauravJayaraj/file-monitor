# file-monitor

- v1
* thoughts
    - while we are monitoring what if someone deletes the file?
        - soln: 
            check for file info(inode)
            using os.stat(), and os.isnotexist(err)
            
            under the hood -> unix-> fstat() system call to get file inode then access() method to check if deleted.

    - remote file monitoring ( inspired by scp )


- shoutout to Mr.Abhishek Jindal for giving me the inspiration for this tool!