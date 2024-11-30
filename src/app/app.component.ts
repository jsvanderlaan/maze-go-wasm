import { CommonModule } from '@angular/common';
import { Component, inject, OnInit } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { ProcessService } from 'src/services/process.service';
import { WasmService } from 'src/services/wasm.service';
import { FileUploadComponent } from './file-upload.component';
import { ResultComponent } from './result.component';
import { SettingsComponent } from './settings.component';
import { ExpanderComponent } from './expander.component';

@Component({
    imports: [CommonModule, ReactiveFormsModule, FileUploadComponent, ResultComponent, SettingsComponent, ExpanderComponent],
    selector: 'app-root',
    templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {
    private readonly wasmService = inject(WasmService);
    readonly processService = inject(ProcessService);
    readonly toggleFileExpander = this.processService.original.asObservable();
    async ngOnInit(): Promise<void> {
        await this.wasmService.loadAndRunGoWasm();
    }
}
