import { BrowserRouter, Routes, Route } from "react-router-dom";
import IndexPage from "./pages";
import ClassPage from "./pages/classes";
import StudyGroupPage from "./pages/studyGroup";
import LogIn from "./pages/login";
import SecuredRoute from "./components/SecuredRoute";
import GlobalMenu from "./components/GlobalMenu";


const Router: React.FC = () => (
    <BrowserRouter>
        <Routes>
            <Route element={<SecuredRoute />}>
                <Route element={<GlobalMenu />}>
                    <Route path="/" element={<IndexPage />} />
                    <Route path="/class" element={<ClassPage />} />
                    <Route path="/studyGroup" element={<StudyGroupPage />} />

                </Route>
            </Route>
            <Route path="/login" element={<LogIn />} />
        </Routes>
    </BrowserRouter>
)

export default Router;