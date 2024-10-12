// server.js

const express = require('express');
const exphbs  = require('express-handlebars');
const path = require('path');

// Initialize Express app
const app = express();

// Set up Handlebars middleware
app.engine('handlebars', exphbs.engine({ defaultLayout: 'main' }));
app.set('view engine', 'handlebars');

// Serve static files from the 'public' directory
app.use(express.static(path.join(__dirname, 'public')));

// Import routes
const indexRoutes = require('./routes/index');
app.use('/', indexRoutes);

// Define the port
const PORT = process.env.PORT || 3000;

// Start the server
app.listen(PORT, () => {
  console.log(`Server started on http://localhost:${PORT}`);
});
