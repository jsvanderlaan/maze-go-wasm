import { NgModule, Pipe, PipeTransform } from '@angular/core';
import { Status, StatusType } from 'src/types/status.type';

@Pipe({
    name: 'loading',
    standalone: true,
})
export class LoadingPipe<T> implements PipeTransform {
    transform(value: Status<T>): boolean {
        return value.status === StatusType.loading;
    }
}

@Pipe({
    name: 'error',
    standalone: true,
})
export class ErrorPipe<T> implements PipeTransform {
    transform(value: Status<T>): boolean {
        return value.status === StatusType.error;
    }
}

@Pipe({
    name: 'done',
    standalone: true,
})
export class DonePipe<T> implements PipeTransform {
    transform(value: Status<T>): T | null {
        if (value.status !== StatusType.done) {
            return null;
        }
        return value.value;
    }
}

@NgModule({
    imports: [LoadingPipe, ErrorPipe, DonePipe],
    exports: [LoadingPipe, ErrorPipe, DonePipe],
})
export class StatusModule {}
