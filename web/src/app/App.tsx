import { BrowserRouter } from 'react-router-dom';
import { Provider } from 'react-redux';

import { store } from './store';
import { Switch } from './routes';

function App() {
    return (
        <Provider store={store}>
            <BrowserRouter>
                <Switch />
            </BrowserRouter>
        </Provider>
    );
}

export default App;
