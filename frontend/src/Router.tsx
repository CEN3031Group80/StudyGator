import { BrowserRouter, Routes, Route } from "react-router-dom";
import IndexPage from "./pages";
import LogIn from "./pages/login";
import SecuredRoute from "./components/SecuredRoute";


const Router: React.FC = () => (
    <BrowserRouter>
        <Routes>
            <Route element={<SecuredRoute />}>
            <Route path="/" element={<IndexPage />} />
            </Route>
            <Route path="/login" element={<LogIn />} />
        </Routes>
    </BrowserRouter>
)

export default Router;