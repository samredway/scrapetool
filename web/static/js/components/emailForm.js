import { sendEmail } from '../services/apiService.js';

export class EmailForm {
    constructor() {
        this.form = document.getElementById('emailForm');
        this.bindEvents();
    }

    bindEvents() {
        this.form.addEventListener('submit', (e) => this.handleSubmit(e));
    }

    async handleSubmit(e) {
        e.preventDefault();
        const formData = new FormData(this.form);
        const email = formData.get('email')?.trim();
        await sendEmail(email);
    }
}
