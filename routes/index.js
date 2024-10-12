import express from 'express';
const router = express.Router();

// Home route
router.get('/', (req, res) => {
  res.render('index');
});

// Get Started route - renders the form
router.get('/get-started', (req, res) => {
  res.render('form');
});

// Handle form submission
router.post('/register', (req, res) => {
  const { email, phone, contacts } = req.body;
  console.log(`User registered: Email - ${email}, Phone - ${phone}, Emergency Contacts - ${contacts}`);
  res.render('video-placeholder', { email, phone, contacts });
});

// Contact Us route - renders the contact form
router.get('/contact', (req, res) => {
  res.render('contact');
});

// Handle contact form submission
router.post('/contact', (req, res) => {
  const { email, message } = req.body;
  console.log(`Contact request received: Email - ${email}, Message - ${message}`);
  res.send('Thank you for reaching out to us. We will get back to you shortly.');
});

export default router;
