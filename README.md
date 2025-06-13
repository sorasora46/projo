# projo
Minimalist project tracker written with React.js + Go Fiber

### Features
- user must authenticate themself to the system
- user should be able perform CRUD ops project(s)
- user should be able perform CRUD ops project's task(s)
- user should be able to view how many tasks left
- user should be able to view all projects in dashboard
- each project in dashboard should have created date, title, DONE task percentage
- user can view percentage of DONE/non-DONE tasks
- task should have status: TODO, DOING, DONE

### TODO [frontend]
 - [ ] create login page
 - [ ] create home page
### TODO [backend]
 - [ ] implement project tasks usecase
 - [ ] create SIT for user endpoint
 - [ ] create E2E for user endpoint
 - [ ] create SIT for project endpoint
 - [ ] create E2E for project endpoint
 - [x] validate all request dtos
 - [x] fix unit test of user endpoint
 - [x] refactor response objs
 - [x] refactor error handling
 - [x] change http method name to constant
 - [x] change http status to constant
 - [x] setup userId in JWT and validate user before sending to next middleware
 - [x] setup github CI
 - [x] create CRUD project endpoint
 - [x] create unit test for user CRUD
 - [x] create login endpoint
 - [x] handle auth middleware
 - [x] create CRUD user endpoint
 - [x] setup GORM (better structured code)
 - [x] refactor init ENV

### UI Resource
Similar to zed UI or these
- https://dribbble.com/shots/25683483-Dashboard-UI
- https://dribbble.com/shots/25155029-Finance-Dashboard-Design
- https://dribbble.com/shots/23355899-Finance-Dashboard-design
- https://dribbble.com/shots/23185186-Finance-Dashboard-Design
