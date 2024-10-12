import express from 'express';
const router = express.Router();

// Home route
router.get('/', (req, res) => {
  res.render('index');
});

export default router;  // Use ES module export
