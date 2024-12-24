export enum StatusType {
    loading,
    done,
    error,
}

export type LoadingStatus = { status: StatusType.loading };
export type ErrorStatus = { status: StatusType.error; error: Error };
export type DoneStatus<T> = { status: StatusType.done; value: T };
export type Status<T> = LoadingStatus | ErrorStatus | DoneStatus<T>;
