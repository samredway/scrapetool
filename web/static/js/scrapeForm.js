import { scrapeUrl } from './apiService.js';

export class ScrapeForm {
    constructor() {
        this.form = document.getElementById('scrapeForm');
        this.resultElement = document.getElementById('scrapeResult');
        this.bindEvents();
    }

    bindEvents() {
        this.form?.addEventListener('submit', (e) => this.handleSubmit(e));
    }

    async handleSubmit(event) {
        event.preventDefault();
        const formData = new FormData(this.form);
        const url = formData.get('url')?.trim();
        const prompt = formData.get('prompt')?.trim();

        if (!url) {
            this.showError('Please enter a URL');
            return;
        }

        try {
            this.showLoading();
            const results = await scrapeUrl(url, prompt);
            this.showResults(results);
        } catch (error) {
            console.error('Error:', error);
            this.showError('An error occurred while scraping the content.');
        }
    }

    showLoading() {
        this.resultElement.innerText = 'Loading...';
    }

    showResults(results) {
        this.resultElement.innerText = results;
    }

    showError(message) {
        this.resultElement.innerText = message;
    }
}
