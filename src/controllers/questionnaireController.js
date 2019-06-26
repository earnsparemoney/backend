const { Questionnaire, User, Participation } = require('../models')
const jwt = require('jsonwebtoken')
const config = require('../config/config')

module.exports = {
  async getAllQuestionnaires (req, res) {
    try {
      var questionnaires = await Questionnaire.findAll({
        include: [{ model: User, as: 'publisher', attributes: ['id', 'username', 'email', 'phone', 'img'] }]
      })
      res.send({
        questionnaires: questionnaires
      })
    } catch (err) {
      res.status(400).send({
        error: 'Some wrong occoured when getting data!'
      })
    }
  },
  async getPublishedQuestionnaires (req, res) {
    try {
      var questionnaires = await Questionnaire.findAll({
        where: {
          publisherId: req.params.id
        },
        include: [{ model: User, as: 'owner', attributes: ['id', 'username', 'email', 'phone', 'img'] }]
      })
      res.send({
        questionnaires: questionnaires
      })
    } catch (err) {
      res.status(400).send({
        error: 'Token should be given, Please check your login state!'
      })
    }
  },
  async addQuestionnaire (req, res) {
    try {
      const token = req.header('Authorization')
      if (!token) {
        return res.status(400).send({
          error: 'token should be given!'
        })
      }
      var result = null
      try {
        result = jwt.verify(token, config.authServiceToken.secretKey)
        if (!result) {
          return res.status(400).send({
            error: 'The token is not valid! Please sign in and try again!'
          })
        }
      } catch (err) {
        return res.status(400).send({
          error: 'Token expired, please login again!'
        })
      }
      await Questionnaire.create({
        title: req.body.title,
        questions: req.body.questions,
        startDate: req.body.startDate,
        endDate: req.body.endDate,
        publisherId: result.id,
        adward: req.body.adward,
        usernum: req.body.usernum
      })

      var questionnaire = await Questionnaire.findOne({
        where: {
          title: req.body.title
        },
        include: [{ model: User, as: 'publisher', attributes: ['id', 'username', 'email', 'phone', 'img'] }]
      })

      res.send({
        questionnaire: questionnaire.toJSON()
      })
    } catch (err) {
      console.log(err)
      res.status(400).send({
        error: err.fields !== undefined ? 'The ' + err.fields[0] + ' has been used!' : 'Error input please check your input'
      })
    }
  },
  async deleteQuestionnaire (req, res) {
    try {
      const token = req.header('Authorization')
      if (!token) {
        return res.status(400).send({
          error: 'token should be given!'
        })
      }
      var result = null
      try {
        result = jwt.verify(token, config.authServiceToken.secretKey)
        if (!result) {
          return res.status(400).send({
            error: 'The token is not valid! Please sign in and try again!'
          })
        }
      } catch (err) {
        return res.status(400).send({
          error: 'Token expired, please login again!'
        })
      }
      var questionnaire = await Questionnaire.findOne({
        where: {
          id: req.params.id
        }
      })
      if (!questionnaire) {
        res.status(400).send({
          error: 'No place is found, please check your request!'
        })
      }
      await questionnaire.destroy()
      res.send({
        info: 'Delete place successfully'
      })
    } catch (err) {
      res.status(400).send({
        error: 'Some error occured when deleting event!'
      })
    }
  },
  async getParticipateQuestionnaire (req, res) {
    try {
      //const token = req.header('Authorization')
      var result = []
      id = req.params.id
      var paticipations = await Participation.findAll({
        where: {
          UserId: id
        }
      })

      if(!paticipations) {
        res.status(400).send({
          error: "user not attend any questionnaire or not such user"
        })
      }

      for (pars in paticipations){
        let qid = participations[pars].id
        let quesOne = await Questionnaire.findOne({
          where: {
            id: qid
          }
        })
        result.push(quesOne)
      }

      res.send({
        questionnaires: JSON.stringify(result)
      })

      
    } catch (error) {
      res.status(400).send({
        error: error
      })
    }
    return null
  },
  async participateQuestionnaire (req, res) {
    try {
      const token = req.header('Authorization')
      if (!token) {
        return res.status(400).send({
          error: 'token should be given!'
        })
      }
      var result = null
      try {
        result = jwt.verify(token, config.authServiceToken.secretKey)
        if (!result) {
          return res.status(400).send({
            error: 'The token is not valid! Please sign in and try again!'
          })
        }
      } catch (err) {
        return res.status(400).send({
          error: 'Token expired, please login again!'
        })
      }

      let uid = result.id
      let qid = req.params.id
      let answer = req.body.answer

      let par = await Participation.findOne({
        where: {
          UserId: uid,
          QuestionnaireId: qid
        }
      })

      if(!par){
        res.status(400).send({
          error: "user already complete this questionnaire"
        })
      }

      await Participation.create({
        UserId: uid,
        QuestionnaireId: qid,
        answer: answer
      })

      res.send({
        info: "success"
      })

    } catch (error) {
      res.status(400).send({
        error: error
      })
    }
    return null
  },
  async getDetil (req, res) {
    try {
      var questionnaire = await Questionnaire.findOne({
        where: {
          id: req.params.id
        },
        include: [{ model: User, as: 'publisher', attributes: ['id', 'username', 'email', 'phone', 'img'] }]
      })
      res.send({
        questionnaire: questionnaire
      })
    } catch (err) {
      res.status(400).send({
        error: 'Some error occured when deleting event!'
      })
    }
  },

  async getAllResults (req, res) {
    const token = req.header('Authorization')
    if (!token) {
      return res.status(400).send({
        error: 'token should be given!'
      })
    }
    var result = null
    try {
      result = jwt.verify(token, config.authServiceToken.secretKey)
      if (!result) {
        return res.status(400).send({
          error: 'The token is not valid! Please sign in and try again!'
        })
      }
    } catch (err) {
      return res.status(400).send({
        error: 'Token expired, please login again!'
      })
    }

    let uid = result.id
    let qid = req.params.qid

    let quesFind = await Questionnaire.findOne({
      where: {
        id: qid,
        publisherId: uid
      }
    })

    if(!quesFind){
      res.status(400).send({
        error: "you don't own this questionnaire"
      })
    }


    let all = await Participation.findAll({
      where: {
        QuestionnaireId: qid
      }
    })

    let allResult = all.join('\n')
    
    res.send({
      result: allResult
    })

  }
}
