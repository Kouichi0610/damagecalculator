// types.ts
// state定義
export interface TodosState {
    title: string;
    todos: Todo[];
}
export interface Todo {
    id: string;
    done: boolean;
    text: string;
}
