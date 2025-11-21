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
    readonly defaultHeight = 50;
    readonly minHeight = 1;
    readonly maxHeight = 1000;
    readonly submitted = signal<boolean>(false);
    private processService = inject(ProcessService);

    readonly form: FormGroup<{
        text: FormControl<string>;
        outline: FormControl<boolean>;
        height: FormControl<number>;
    }> = new FormGroup({
        text: new FormControl('Amazing', {
            nonNullable: true,
            validators: [Validators.required],
        }),
        height: new FormControl(this.defaultHeight, {
            nonNullable: true,
            validators: [Validators.min(this.minHeight), Validators.max(this.maxHeight), Validators.required],
        }),
        outline: new FormControl(false, {
            nonNullable: true,
        }),
    });

    async onSubmit(): Promise<void> {
        this.submitted.set(true);
        if (this.form.valid) {
            this.processService.sourceText.next(this.form.value as any);
        }
    }
}
