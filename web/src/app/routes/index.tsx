import { Navbar } from '@components/common/Navbar';
import { Footer } from '@components/common/Footer';
import { lazy, useEffect } from 'react';
import { Route, Routes, useLocation, useNavigate } from 'react-router-dom';
import { UserEdit } from '@components/User/Edit';

const Home = lazy(() => import('@/app/views/Home'));
const Registration = lazy(() => import('@/app/views/Auth/Registration'));
const Login = lazy(() => import('@/app/views/Auth/Login'));
const User = lazy(() => import('@/app/views/User'));
/**
 * ComponentRoutes describes location mapping with components.
 */
export class ComponentRoutes {
    constructor(
        public path: string,
        public element: JSX.Element,
        public children?: ComponentRoutes[]
    ) { }

    /** with is method that creates child sub routes path */
    public with(
        child: ComponentRoutes,
        parrent: ComponentRoutes
    ): ComponentRoutes {
        child.path = `${parrent.path}/${child.path}`;

        return this;
    }

    /** addChildren is method that adds children components to component */
    public addChildren(children: ComponentRoutes[]): ComponentRoutes {
        this.children = children.map((child: ComponentRoutes) =>
            child.with(child, this)
        );

        return this;
    }
}

/**
 * RoutesConfig contains information about all routes and subroutes.
 */
export class RoutesConfig {
    public static Home: ComponentRoutes = new ComponentRoutes(
        '/',
        <Home />
    );

    public static Registration: ComponentRoutes = new ComponentRoutes(
        '/',
        <Registration />
    );
    public static Login: ComponentRoutes = new ComponentRoutes(
        '/',
        <Login />
    );
    public static User: ComponentRoutes = new ComponentRoutes(
        '/user/:id',
        <User />
    );
    public static UserEdit: ComponentRoutes = new ComponentRoutes(
        '/user/:id/edit',
        <UserEdit />
    );
    /** Routes is an array of logical router components */
    public static routes: ComponentRoutes[] = [
        RoutesConfig.Home,
        RoutesConfig.Registration,
        RoutesConfig.Login,
        RoutesConfig.User,
        RoutesConfig.UserEdit,
    ];
}

export class AuthRoutesConfig {
    public static Login: ComponentRoutes = new ComponentRoutes(
        '/login',
        <Login/>,
    );
    public static Registration: ComponentRoutes = new ComponentRoutes(
        '/registration',
        <Registration/>,
    );
    /** Routes is an array of logical router components */
    public static routes: ComponentRoutes[] = [
        AuthRoutesConfig.Login,
        AuthRoutesConfig.Registration,
    ];
}

export const Switch = () => {

     const location = useLocation();
    const navigate = useNavigate();

    const isLoggedin = true;

    useEffect(() => {
        if (!isLoggedin && location.pathname !== AuthRoutesConfig.Login.path) {
            console.log(location.pathname )
            navigate(AuthRoutesConfig.Registration.path);
        } else if (!isLoggedin && location.pathname === AuthRoutesConfig.Login.path) {
            console.log('hello')
            navigate(AuthRoutesConfig.Login.path);
        } else {
            navigate(RoutesConfig.Home.path);
        }
    }, []);

    return (
        <>
            {!isLoggedin ?
                <Routes>
                    {AuthRoutesConfig.routes.map(
                        (route: ComponentRoutes, index: number) =>
                            <Route
                                key={index}
                                path={route.path}
                                element={route.element}
                            />
                    )}
                </Routes>
                : <div>
                    <Navbar />
                    <div className="page">
                        <Routes>
                            {RoutesConfig.routes.map(
                                (route: ComponentRoutes, index: number) =>
                                    <Route
                                        key={index}
                                        path={route.path}
                                        element={route.element}
                                    />
                            )}
                        </Routes>
                    </div>
                    <Footer />
                </div>
            }
        </>
        );
};
