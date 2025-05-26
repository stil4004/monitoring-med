# QA service

Main service functionality is to collect and research questions and answers. Connected with two main logics: solved and unresolved questions.

## Static questions
\
Table and entity that provides functions to add/delete and collect static questions.

### /questions/all POST

Route for collecting all questions ordered in DESC time \
*Request*: 
```
{
    "question": "smquestion", 
    "answer": "smanswer",
    "project_type_id": 1
}
```

*Response*:
```
{
    "id": "25"
}
```

### /questions/add POST

Route for adding new question  
*Request*: 
```
{
    "question": "smquestion", 
    "answer": "smanswer",
    "project_type_id": 1
}
```

*Response*:
```
{
    "id": "25"
}
```
### /questions/:id GET

Route for getting question by id  
*Request*: 
```
"id" in path (question_id)
```

*Response*:
```
{
    "id": "25"
}
```

### /questions/:id GET

Route for getting question by id  
*Request*: 
```
"id" in path (question_id)
```

*Response*:
```
{
    "id": "25"
}
```

### /questions/delete DELETE

Route for adding new question  
*Request*: 
```
{
    "question_id": 17
}
```

*Response*:
```
{
    "status": "success"
}
```