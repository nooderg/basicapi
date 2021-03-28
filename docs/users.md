# Articles API reference


## Routes

<br/>

### Authenticate an user
#### (POST) `/login`

<br/>
 
Parameters: [LoginForm](#LoginForm)  
Returns: [LoggedUser](#LoggedUser) `(Object)`  
Authentication needed: `false`

<br/>

### Register a new user
#### (POST) `/register`

<br/>
 
Parameters: [UserForm](#UserForm)  
Returns: `true` or `false` `(Bool)`  
Authentication needed: `false`

<br/>

### Get logged in user
#### (GET) `/register`

<br/>
 
Parameters: JWT token `(String)`  
Returns: [LoggedUser](#LoggedUser) `(Object)`  
Authentication needed: `true` 

<br/>

### Edit currently logged in user
#### (POST) `/register`

<br/>
 
Parameters: JWT token `(String)`  
Returns: [LoggedUser](#LoggedUser) `(Object)`  
Authentication needed: `true` 

<br/>

## Models

<br/>

### LoggedUser

Type: `Object`

Parameters:

• `ID` (String) - User ID  
• `username` (String) - Username  
• `firstname` (String) - First name  
• `lastname` (String) - Last name   
• `dob` (String) - Date of birth  
• `city` (String) - City  
• `jwt` (String) - JWT authentication token

<br/>

### UserForm

Type: `Object`

Parameters:

• `username` (String) - Username choice  
• `password` (String) - Password choice  
• `confirm_password` (String) - Password verification  
• `firstname` (String) - First name
• `lastname` (String) - Last name
• `dob` (String) - Date of birth
• `city` (String) - City

<br/>

### LoginForm

Type: `Object`  

Parameters:

• `username` (String)  
• `password` (String)  