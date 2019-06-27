const authController = require('./controllers/authController')
const authControllerPolicy = require('./policies/authControllerPolicy')
const taskController = require('./controllers/taskController')
const questionnaireController = require('./controllers/questionnaireController')
const uploader = require('./utils/uploader')

module.exports = (app) => {
  /***
   * Auth Part
   */
  app.get('/', (req, res) => {
    res.send('index')
  })
  app.post('/user',
    uploader.userImg.single('image'),
    authControllerPolicy.register,
    authController.register)
  app.post('/user/login',
    authController.login)
  app.put('/user',
    uploader.userImg.single('image'),
    authController.updateUser)
  app.get('/usericon/:username',
    authController.getIcon)
  /***
   * Task Part
   */
  app.get('/tasks',
    taskController.getAllTasks)
  app.post('/tasks',
    taskController.addTask)
  app.delete('/tasks/:id',
    taskController.deleteTask)
  app.get('/tasks/user/:id',
    taskController.getAllPublishedTasks)
  app.get('/tasks/user/:id/finish',
    taskController.getAllFinishedTasksParticipatesIn)
  app.get('/tasks/user/:id/running',
    taskController.getAllRunningTasksParticipatesIn)
  app.post('/task/:id/participate',
    taskController.participateTask)
  app.post('/task/:id/finish',
    taskController.finishTask)
  app.get('/task/:id',
    taskController.getDetail)
  /***
   * Questionnaire Part
   */
  app.get('/questionnaires',
    questionnaireController.getAllQuestionnaires)
  app.post('/questionnaires',
    questionnaireController.addQuestionnaire)
  app.get('/questionnaires/participates/user/:id',
    questionnaireController.getParticipateQuestionnaire)
  app.get('/questionnaire/:id',
    questionnaireController.getDetil)
  app.delete('/questionnaires/:id',
    questionnaireController.deleteQuestionnaire)
  app.get('/questionnaires/user',
    questionnaireController.getPublishedQuestionnaires)
  app.post('/questionnaires/:id',
    questionnaireController.participateQuestionnaire)
  app.get('/results/:qid',
    questionnaireController.getAllResults)
}
