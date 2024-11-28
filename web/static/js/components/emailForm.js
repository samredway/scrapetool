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
        
        try {
            await sendEmail(email);
            this.form.reset();
            this.showSuccess();
        } catch (error) {
            console.error('Failed to register email:', error);
        }
    }

    showSuccess() {
        // Remove any existing success message
        const existingSuccess = this.form.querySelector('.alert-success');
        if (existingSuccess) {
            existingSuccess.remove();
        }

        // Create and show new success message
        const successDiv = document.createElement('div');
        successDiv.className = 'alert alert-success mt-3';
        successDiv.textContent = 'Thanks! We\'ll keep you updated about early access and launch details.';
        this.form.appendChild(successDiv);

        // Remove after 5 seconds
        setTimeout(() => {
            successDiv.remove();
        }, 5000);
    }
}
