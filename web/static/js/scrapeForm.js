import { scrapeUrl } from './apiService.js';

export class ScrapeForm {
    constructor() {
        this.form = document.getElementById('scrapeForm');
        this.resultElement = document.getElementById('scrapeResult');
        this.responseStructureSection = document.getElementById('responseStructureSection');
        this.specifyResponseStructureCheckbox = document.getElementById('specifyResponseStructure');
        
        // Initialize event listeners
        this.bindEvents();
    }

    bindEvents() {
        this.form.addEventListener('submit', (e) => this.handleSubmit(e));
        this.specifyResponseStructureCheckbox.addEventListener('change', () => this.toggleResponseStructureSection());
    }

    toggleResponseStructureSection() {
        if (this.specifyResponseStructureCheckbox.checked) {
            this.responseStructureSection.style.display = 'block';
        } else {
            this.responseStructureSection.style.display = 'none';
        }
    }

    async handleSubmit(event) {
        event.preventDefault();
        const formData = new FormData(this.form);
        const url = formData.get('url')?.trim();
        const prompt = formData.get('prompt')?.trim();
        const responseStructure = formData.get('responseStructure')?.trim();

        if (!url) {
            this.showError('Please enter a URL');
            return;
        }

        if (prompt === '') {
            this.showError('Please enter a prompt');
            return;
        }

        this.showLoading();

        try {
            const results = await scrapeUrl(url, prompt, responseStructure);
            this.showResults(results);
        } catch (error) {
            this.showError(`${error}`);
            return;
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
