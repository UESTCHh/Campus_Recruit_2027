# 这个实验的目标不是写业务，而是亲手理解 async/await 的执行模型。
import asyncio
from time import perf_counter


async def fake_request(name: str) -> str:
    print(f"{name} started")
    await asyncio.sleep(1)
    print(f"{name} finished")
    return name


async def run_sequential() -> None:
    start = perf_counter()

    # 依次 await 三次
    await fake_request("A")
    await fake_request("B")
    await fake_request("C")

    elapsed = perf_counter() - start
    print(f"sequential: {elapsed:.2f}s")


async def run_concurrent() -> None:
    start = perf_counter()

    # 使用 asyncio.gather 并发运行三个任务
    await asyncio.gather(fake_request("A"), fake_request("B"), fake_request("C"))

    elapsed = perf_counter() - start
    print(f"concurrent: {elapsed:.2f}s")


async def main() -> None:
    await run_sequential()
    await run_concurrent()


if __name__ == "__main__":
    asyncio.run(main())
