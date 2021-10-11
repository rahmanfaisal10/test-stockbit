# test-stockbit

## Answer Question Technical Test stockbit

### Answer Number One 

Jawaban dari pertanyaan no 1 anda bisa dapatkan di folder: 
    
    number_one/query_answer

atau anda bisa lihat query nya seperti dibawah ini:
    
    SELECT 
        u.ID, 
        u.UserName, 
        u2.UserName as ParentUserName 
    FROM `USER` u 
    LEFT JOIN `USER` u2 ON u.Parent = u2.ID;

### Answer Number Two

Anda dapat menjalankan jawaban no 2 dengan:

apabila anda hanya menjalankan transport http saja pakai:

    make start-http

Router untuk menjalankan http response adalah:
    endpoint http yang digunakan adalah:
        http://localhost:8080/api/v1/omdbapi/detail-movie-by-id?i=tt4853102 // endpoint detail by id dengan id = tt4853102
        http://localhost:8080/api/v1/omdbapi/list-movie?s=Batman&pages=1 // endpoint list movie dengan searchword Batman dan pagination 1

apabila anda hanya menjalankan transport http saja pakai:

    make start-grpc
### Answer Number Three

Jawaban dari pertanyaan no 3 anda bisa dapatkan di folder:

    number_three/answe_three.go

atau anda juga bisa lihat methodnya seperti dibawah ini:

    func FindFirstStringInBracket(str string) string {
        indexFirstBracketFound := strings.Index(str, "(")
        if indexFirstBracketFound >= 0 {
            indexClosingBracketFound := strings.Index(str, ")")
            switch {
            case indexFirstBracketFound > indexClosingBracketFound:
                return ""
            case indexClosingBracketFound >= 0:
                return str[indexFirstBracketFound+1 : indexClosingBracketFound]
            default:
                return ""
            }
        }

        return ""
    }

## Answer Number Four

Anda Dapat Menjalankan Jawaban no 4 dengan:

    make start-four

atau anda bisa lihat service yang bisa jalan di file:

    Makefile