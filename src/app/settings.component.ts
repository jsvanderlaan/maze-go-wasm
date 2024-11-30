import { CommonModule } from '@angular/common';
import { Component, inject } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Observable } from 'rxjs';
import { ProcessService } from 'src/services/process.service';
import { ExpanderComponent } from './expander.component';

@Component({
    imports: [ReactiveFormsModule, CommonModule, ExpanderComponent],
    templateUrl: './settings.component.html',
    selector: 'app-settings',
})
export class SettingsComponent {
    static readonly defaultSettings: Settings = { size: 100 };

    private readonly processService = inject(ProcessService);

    readonly defaultSettings = SettingsComponent.defaultSettings;
    readonly currentSettings$: Observable<Settings> = this.processService.settings;
    readonly toggle$ = this.processService.settings.asObservable();

    readonly sizeMin: number = 10;
    readonly sizeMax: number = 400;

    readonly form = new FormGroup({
        size: new FormControl(SettingsComponent.defaultSettings.size, {
            nonNullable: true,
            validators: [Validators.min(this.sizeMin), Validators.max(this.sizeMax)],
        }),
    });

    submit(): void {
        const settings = this.form.getRawValue();
        this.processService.settings.next(settings);
    }
}

export interface Settings {
    size: number;
}
