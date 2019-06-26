const multer = require('multer')

const quetionnaireStorage = multer.diskStorage({
  destination: 'public/images/questionnaireImage/',
  filename: function (req, file, cb) {
    var fileformat = (file.originalname).split('.')
    cb(null, fileformat[0] + '-' + Date.now() + '.' + fileformat[fileformat.length - 1])
  }
})

const taskStorage = multer.diskStorage({
  destination: 'public/images/taskImage/',
  filename: function (req, file, cb) {
    var fileformat = (file.originalname).split('.')
    cb(null, fileformat[0] + '-' + Date.now() + '.' + fileformat[fileformat.length - 1])
  }
})

const userStorage = multer.diskStorage({
  destination: 'public/images/userImage/',
  filename: function (req, file, cb) {
    var fileformat = (file.originalname).split('.')
    cb(null, fileformat[0] + '-' + Date.now() + '.' + fileformat[fileformat.length - 1])
  }
})

module.exports = {
  quetionnaireImg: multer({
    storage: quetionnaireStorage
  }),
  taskImg: multer({
    storage: taskStorage
  }),
  userImg: multer({
    storage: userStorage
  })
}
