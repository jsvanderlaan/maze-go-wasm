import { CommonModule } from '@angular/common';
import { Component, inject } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { merge } from 'rxjs';
import { StatusModule } from 'src/pipes/status.pipe';
import { ProcessService } from 'src/services/process.service';
import { ResultComponent } from './result/result.component';
import { ExpanderComponent } from './shared/expander.component';
import { SpinnerComponent } from './shared/spinner.component';
import { SourceComponent } from './source/source.component';

@Component({
    imports: [
        CommonModule,
        ReactiveFormsModule,
        SourceComponent,
        ResultComponent,
        ExpanderComponent,
        StatusModule,
        SpinnerComponent,
    ],
    selector: 'app-root',
    templateUrl: './app.component.html',
})
export class AppComponent {
    readonly processService = inject(ProcessService);
    readonly toggleFileExpander = merge(
        this.processService.sourceImage.asObservable(),
        this.processService.sourceText.asObservable()
    );
}
