import express from 'express';
import path from 'path';
import { fileURLToPath } from 'url';
import fs from 'fs';
import hbs from 'hbs';
import { exec } from 'child_process';

const app = express();

app.set('view engine', 'hbs');

const __dirname = path.dirname(fileURLToPath(import.meta.url));
app.use(express.urlencoded({ extended: true }));

// Set the views and partials directories
app.set('views', path.join(__dirname, 'views'));
hbs.registerPartials(path.join(__dirname, 'views', 'partials'));

app.use(express.static(path.join(__dirname, 'public')));

// Import routes
import indexRoutes from './routes/index.js';
app.use('/', indexRoutes);

// Add route to serve the features page
app.get('/features', (req, res) => {
  res.render('features');
});

// Add route to serve the contact page
app.get('/contact', (req, res) => {
  res.render('contact');
});

// Add route to start the video feed
app.get('/start-video-feed', (req, res) => {
  exec('python3 ../HackHarvard/machine_learning/app2.py', (error, stdout, stderr) => {
    if (error) {
      console.error(`Error executing Python script: ${error}`);
      return res.status(500).json({ success: false, message: 'Error starting video feed' });
    }
    console.log(`Python script output: ${stdout}`);
    res.json({ success: true, message: 'Video feed started successfully' });
  });
});

// Define the port
const PORT = process.env.PORT || 3000;

// Start the server
app.listen(PORT, () => {
  console.log(`Server started on http://localhost:${PORT}`);
});
