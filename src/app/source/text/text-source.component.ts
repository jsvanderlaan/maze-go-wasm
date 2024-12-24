import { Component, inject, signal } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { ProcessService } from 'src/services/process.service';

@Component({
    standalone: true,
    imports: [ReactiveFormsModule],
    selector: 'app-text-source',
    templateUrl: './text-source.component.html',
})
export class TextSourceComponent {
    readonly submitted = signal<boolean>(false);
    private processService = inject(ProcessService);
    readonly form: FormGroup<{ text: FormControl<string | null>; outline: FormControl<boolean | null> }> = new FormGroup({
        text: new FormControl('', Validators.required),
        outline: new FormControl(false),
    });

    async onSubmit(): Promise<void> {
        this.submitted.set(true);
        if (this.form.valid && this.form.value.text) {
            this.processService.sourceText.next(this.form.value as any);
        }
    }
}
