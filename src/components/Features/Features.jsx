import React from "react";

export default function Features() {
  return (
    <>
      <div>
        <div className="all mt-10">
          <div className="header">
            <div className="flex justify-between">
              <img
                src="./images/mainweb/3D/Galaxy 1.png"
                alt=""
                className="w-[40px]"
              />
              <div className="flex flex-col justify-center text-center">
                <h1 className="text-2xl font-Ray-Black">
                  ویژگی های شرکت{" "}
                  <span className="text-main-blue-web">لاول کد</span>
                </h1>
                <span className="text-main-green-web text-sm">مختص شماست</span>
              </div>

              <img
                src="./images/mainweb/3D/Galaxy 2.png"
                alt=""
                className="w-[50px]"
              />
            </div>
          </div>
          <div className="allMain text-center mr-[7px] md:mr-[100px] lg:mr-[210px] xl:mr-[0]  ">
            <div className="main xl:grid md:grid-cols-12 mt-10 ">
              <div className="right md:col-span-6 md:col-start-2 md:col-end-7 ">
                <div className="box md:w-[550px] md:h-[540px] w-[420px] h-[500px] bg-[#FFEBD3]   rounded-3xl">
                  <div className="img w-[300px] h-[250px]  md:mr-[120px] mr-[60px]">
                    <img
                      src="./images/mainweb/3D/Sec4/5.png"
                      alt=""
                      className=""
                    />
                  </div>

                  <div className="box2 bg-white md:w-[450px] md:h-[230px] w-[380px] h-[200px]  md:mr-[50px] mr-[20px]  rounded-3xl flex flex-col items-center gap-6 mt-[30px]">
                    <div className="mt-4">
                      <h1 className="font-Ray-Black">
                        درصد رضایت و آمار سایت{" "}
                        <span className="text-main-blue-web">لاول کد</span> از
                        زمان ایجاد آن
                      </h1>
                    </div>

                    <div className="font-Ray-Bold">
                      <span>97.6 درصد</span>
                      <span className="mr-[130px]">نرخ رضایت مشتریان </span>
                      <hr className="w-[310px]" />
                    </div>

                    <div className="font-Ray-Bold">
                      <span>3000+</span>
                      <span className="mr-[145px]">مشتری از سراسر ایران</span>
                      <hr className="w-[310px]" />
                    </div>

                    <div className="font-Ray-Bold">
                      <span>50000+</span>
                      <span className="mr-[110px]">
                        کاربران فعال محصولات ما
                      </span>
                    </div>
                  </div>
                </div>
              </div>

              <div className="left md:col-start-7 md:col-end-12 md:col-span-6  mt-[50px] md:mt-[50px] xl:mt-[0]">
                <div className="flex flex-wrap gap-5">
                  <div className="bg-white  md:w-[270px] md:h-[260px] w-[200px] h-[200px] text-center rounded-2xl">
                    <img
                      src="./images/mainweb/3D/Sec4/1.png"
                      alt=""
                      className="md:w-[200px] w-[140px] md:mr-10 mr-[25px]"
                    />
                    <h1 className="font-Ray-ExtraBold md:text-xl text-lg">
                      بهره مندی از متخصیصن
                    </h1>
                  </div>

                  <div className="bg-white  md:w-[270px] md:h-[260px] w-[200px] h-[200px] text-center rounded-2xl">
                    <img
                      src="./images/mainweb/3D/Sec4/2.png"
                      alt=""
                      className="md:w-[200px] w-[140px] md:mr-10 mr-[25px]"
                    />
                    <h1 className="font-Ray-ExtraBold md:text-xl text-lg">
                      امنیت بالا
                    </h1>
                  </div>

                  <div className="bg-white  md:w-[270px] md:h-[260px] w-[200px] h-[200px] text-center rounded-2xl">
                    <img
                      src="./images/mainweb/3D/Sec4/3.png"
                      alt=""
                      className="md:w-[200px] w-[140px] md:mr-10 mr-[25px]"
                    />
                    <h1 className="font-Ray-ExtraBold md:text-xl text-lg">
                      پشتیبانی سریع
                    </h1>
                  </div>

                  <div className="bg-white  md:w-[270px] md:h-[260px] w-[200px] h-[200px] text-center rounded-2xl">
                    <img
                      src="./images/mainweb/3D/Sec4/4.png"
                      alt=""
                      className="md:w-[200px] w-[140px] md:mr-10 mr-[25px]"
                    />
                    <h1 className="font-Ray-ExtraBold md:text-xl text-lg">
                      متعدد به زمان
                    </h1>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
