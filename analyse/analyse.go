package analyse 


import (
  "context"
  "database/sql"
  "os"
  "strings"
  "fmt"
  gogpt "github.com/sashabaranov/go-gpt3"
)


func Sentiment(db *sql.DB) (bool, error){
	
	gpt := gogpt.NewClient(os.Getenv("OPENAPI"))
	ctx := context.Background()
 
     
    sqlStatement := `SELECT id, prompt, text 
                      FROM my_schema.analyser
                    `


    rows, err := db.Query(sqlStatement)
       
    if err != nil {
       return false, err
    }

    defer rows.Close()

    for rows.Next(){

       var id int 
       var prompt string
       var text string
       

       if err := rows.Scan(&id, &prompt, &text); err != nil {
         return false, err
       }

       promptGpt := fmt.Sprintf("%s %s", prompt, text)

       req := gogpt.CompletionRequest{
                Model:       gogpt.GPT3TextDavinci003,
                MaxTokens:   30,
                Temperature: 0.2,
                Prompt:      promptGpt,
        }
        
        
        resp, err := gpt.CreateCompletion(ctx, req)

        if err != nil {
           return false, err
        }

        sqlStatement := `UPDATE my_schema.analyser 
                         SET response = $2
                         WHERE id = $1;`

        fmt.Println(text, resp.Choices[0].Text, "\n")   

        _, err = db.Exec(sqlStatement, id, strings.ReplaceAll(resp.Choices[0].Text, "\n", ""))
        
        if err != nil {
           return false, err
        }


    }

    return true, nil

}
