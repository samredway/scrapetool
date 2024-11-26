import { EmailForm } from './components/emailForm.js';
import { ScrapeForm } from './components/scrapeForm.js';

// Initialize when DOM is ready
document.addEventListener('DOMContentLoaded', () => {
    new ScrapeForm();
    new EmailForm();
});
