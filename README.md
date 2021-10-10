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

### Answer Number Three

Jawaban dari pertanyaan no 3 anda bisa dapatkan di folder:

    number_three/answe_three.go

atau anda juga bisa lihat methodnya seperti dibawah ini:

    ``
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
    ``