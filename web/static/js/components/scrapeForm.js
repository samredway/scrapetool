import { scrapeUrl } from '../services/apiService.js';

const defaultResponseStructure = `{
    "type": "object",
    "properties": {
        "data": {
            "type": "array",
            "items": {"type": "string"}
        }
    },
    "additionalProperties": false,
    "required": ["data"]
}`;

export class ScrapeForm {
    constructor() {
        this.form = document.getElementById('scrapeForm');
        this.resultElement = document.getElementById('scrapeResult');
        this.responseStructureSection = document.getElementById('responseStructureSection');
        this.specifyResponseStructureCheckbox = document.getElementById('specifyResponseStructure');
        this.responseStructureInput = document.querySelector('textarea[name="responseStructure"]');
        this.responseStructureInput.value = defaultResponseStructure;
        this.bindEvents();
    }

    bindEvents() {
        this.form.addEventListener('submit', (e) => this.handleSubmit(e));
        this.specifyResponseStructureCheckbox.addEventListener('change', () => this.toggleResponseStructureSection());
        this.responseStructureInput.addEventListener('keydown', (e) => this.handleResponseStructureTab(e));
    }

    handleResponseStructureTab(e) {
        if (e.key === 'Tab') {
            e.preventDefault();

            // Get cursor position
            const start = e.target.selectionStart;
            const end = e.target.selectionEnd;

            // Insert 4 spaces
            e.target.value = e.target.value.substring(0, start) + "    " + e.target.value.substring(end);

            // Put cursor after the inserted spaces
            e.target.selectionStart = e.target.selectionEnd = start + 4;
        }
    }

    toggleResponseStructureSection() {
        if (this.specifyResponseStructureCheckbox.checked) {
            this.responseStructureSection.style.display = 'block';
        } else {
            this.responseStructureSection.style.display = 'none';
            this.responseStructureInput.value = defaultResponseStructure;
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
