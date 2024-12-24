import { Observable } from 'rxjs';
import { Status, StatusType } from 'src/types/status.type';

export function inspectStatus<T>(source: Observable<T>): Observable<Status<T>> {
    return new Observable(subscriber => {
        subscriber.next({ status: StatusType.loading });
        source.subscribe({
            next: value => subscriber.next({ status: StatusType.done, value }),
            error: err => subscriber.next({ status: StatusType.error, error: err }),
            complete: () => subscriber.complete(),
        });
    });
}
