import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { map, Observable } from 'rxjs';
import { ImgHelper } from 'src/helpers/img.helper';
import { ProcessService } from 'src/services/process.service';

@Component({
    imports: [CommonModule],
    selector: 'app-result',
    templateUrl: './result.component.html'
})
export class ResultComponent {
    result$: Observable<string>;
    constructor(processService: ProcessService) {
        this.result$ = processService.output.pipe(map(arr => ImgHelper.toUrl(arr)));
    }
}
