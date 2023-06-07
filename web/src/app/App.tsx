import { BrowserRouter } from 'react-router-dom';
import { Provider } from 'react-redux';
import { Suspense } from 'react';

import { store } from '@/app/store';
import { Switch } from '@/app/routes';

function App() {
    return (
        <Suspense fallback={'...Loading'}>
            <Provider store={store}>
                <BrowserRouter>
                    <Switch />
                </BrowserRouter>
            </Provider>
        </Suspense>
    );
}

export default App;
