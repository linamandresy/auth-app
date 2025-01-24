import { Routes } from '@angular/router';
import { AuthComponent } from './pages/auth/auth.component';
import { LoginComponent } from './pages/auth/login/login.component';
import path from 'path';

export const routes: Routes = [
    {
        path: '',
        redirectTo: '/auth',
        pathMatch: 'full',
    },{
        path: 'auth',
        component: AuthComponent,
        children: [
            {
                path:'',
                redirectTo: 'login',
                pathMatch:'prefix'
            },{
                path: 'login',
                component: LoginComponent
            }

        ]
    }];
