# test-stockbit

## Answer Question Technical Test stockbit

### Answer Number One 

    Jawaban dari pertanyaan no 1 anda bisa dapatkan di dalam folder number_one dengan penamaan file *query_answer.sql*
    atau anda bisa lihat query nya seperti dibawah ini
    
    ```bash
    SELECT 
        u.ID, 
        u.UserName, 
        u2.UserName as ParentUserName 
    FROM `USER` u 
    LEFT JOIN `USER` u2 ON u.Parent = u2.ID;
    ```