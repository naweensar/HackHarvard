import express from 'express';
import path from 'path';
import { fileURLToPath } from 'url';
import fs from 'fs';
import hbs from 'hbs';

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

// Define the port
const PORT = process.env.PORT || 3000;

// Start the server
app.listen(PORT, () => {
  console.log(`Server started on http://localhost:${PORT}`);
});
