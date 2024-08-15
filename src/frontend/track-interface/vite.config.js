// vite.config.js

import { defineConfig } from 'vite';
import { resolve } from 'path'; // Import the 'resolve' function from Node.js

export default defineConfig({
    // ... other config options ...
    // Configure the server to handle custom routes
    server: {
        proxy: {
            '/apis': {
                //target: 'https://testom2.la0.uk',
                target: 'http://localhost:6051',
                changeOrigin: true,
                withCredentials: true,
            }
        },
    },
});