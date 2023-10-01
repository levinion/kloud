import { createSignal, lazy } from "solid-js";
import { invoke } from "@tauri-apps/api/tauri";
const SideBar = lazy(() => import("./components/SideBar"))
const Header = lazy(() => import("./components/Header"))
const Main=lazy(()=>import("./components/Main"))
function App() {
  let input: any;
  return (
    <div class="flex w-screen min-h-screen">
      {/* sidebar */}
      <div class="bg-gray-500 w-14">
        <SideBar></SideBar>
      </div>
      <div class=" bg-pink-400 flex-grow">
        {/* right */}
        <div class="flex flex-col w-full h-full">
          {/* header */}
          <div class="bg-yellow-200 h-14 w-full">
            <Header></Header>
          </div>
          {/* main */}
          <div class="bg-orange-300 w-full flex-grow">
            <Main></Main>
          </div>
          {/* footer */}
          <div class="bg-purple-400 h-14 w-full">
            {/* <Footer></Footer> */}
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
