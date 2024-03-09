import { BrowserRouter, Routes, Route } from "react-router-dom";
import IndexPage from "./pages";


const Router: React.FC = () => (
    <BrowserRouter>
        <Routes>
            <Route path="/" element={<IndexPage />} />
        </Routes>
    </BrowserRouter>
)

export default Router;