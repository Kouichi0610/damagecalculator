export interface User {
    firstName: string;
    lastName: string;
    email: string;
    phone?: string;
}

export interface ProfileState {
    title: string;
    user?: User;
    error: boolean;
}
