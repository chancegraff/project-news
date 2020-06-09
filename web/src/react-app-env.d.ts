/// <reference types="react-scripts" />

declare module '*.scss' {
    const content: {[className: string]: string};
    export default content;
}

interface IArticle {
    ID: number;
    Title: string;
    URL: string;
    Thumbnail: string;
    PublishedAt: string;
}

interface IUser {
    ID: number;
    Email: string;
}

interface Credentials {
    email: string;
    password: string;
}

interface IRegisterForm extends Credentials {}

interface ILoginForm extends Credentials {}
