from abc import ABC, abstractmethod

# Animal 规定：
# 所有具体动物都必须提供 speak() 方法，并返回字符串。


# 它不负责决定猫怎么叫、狗怎么叫，只负责定义统一规则。
# 对应到你的 AI 项目：
# Animal  ≈ BaseLLM
# speak() ≈ generate()
class Animal(ABC):
    """所有动物实现都必须遵守的统一接口"""

    @abstractmethod
    def speak(self) -> str:
        """返回动物发出的声音。"""


# Cat 和 Dog 都遵守 Animal 的接口，但各自行为不同。
# 对应到 AI 项目：
# Cat         ≈ MockLLM
# Dog         ≈ UppercaseLLM
# 未来其他类   ≈ OpenAILLM、DeepSeekLLM
# 这些实现都可以被当作 Animal 使用。
class Cat(Animal):
    def speak(self) -> str:
        return "喵"


class Dog(Animal):
    def speak(self) -> str:
        return "汪"


# 这个函数就是工厂。
# 输入：
# cat
# 返回：
# Cat()
# 输入：
# dog
# 返回：
# Dog()
# Factory 的职责是：
# 根据条件决定创建哪个具体对象。


# 返回Animal
# 这表示使用者只需要知道：
# animal.speak()
# 不需要关心它究竟是猫还是狗。
# 对应到后续的 LLM Factory：
# def create_llm(provider: str) -> BaseLLM:
#     ...
# 它可能返回：
# MockLLM
# UppercaseLLM
# OpenAILLM
# DeepSeekLLM
# 但调用者只按 BaseLLM 的统一接口使用它。
def create_animal(animal_type: str) -> Animal:
    """根据名称创建对应的 Animal 实现。"""

    # 标准化输入
    normalized_type = animal_type.strip().lower()

    if normalized_type == "cat":
        return Cat()

    if normalized_type == "dog":
        return Dog()

    # Factory 不知道如何创建，就明确抛出：
    # ValueError: Unsupported animal type: bird
    # 这里不能偷偷返回 Cat()，因为错误配置应该清晰失败。
    # 这就是后面要学习的：
    # Fail Fast
    # 即尽早、明确地暴露错误。
    raise ValueError(f"不支持的动物类型: {animal_type}")


# 没有 Factory 会怎样？
# 假设使用者自己负责创建对象：
# animal_type = "cat"
# if animal_type == "cat":
#     animal = Cat()
# elif animal_type == "dog":
#     animal = Dog()
# else:
#     raise ValueError("unsupported animal")
# 如果很多地方都需要创建动物，每个地方都要重复这一段判断：
# 模块 A：判断 cat/dog
# 模块 B：判断 cat/dog
# 模块 C：判断 cat/dog
# 以后增加：
# Bird
# 所有创建对象的位置都可能需要修改。
# 使用 Factory 后：
# animal = create_animal(animal_type)
# 选择逻辑集中在：
# create_animal()
# 使用者只负责：
# animal.speak()
# 职责被分开：
# Factory：决定创建谁
# 调用者：使用创建好的对象
def main() -> None:
    # animal = Animal() #Animal 只定义接口，没有提供完整实现，不能直接实例化。
    cat = create_animal("cat")
    dog = create_animal("dog")

    print(f"cat object: {type(cat).__name__}")
    print(f"cat says: {cat.speak()}")

    print(f"dog object: {type(dog).__name__}")
    print(f"dog says: {dog.speak()}")

    another_cat = create_animal("  CAT  ")
    print(f"another cat says: {another_cat.speak()}")

    try:
        create_animal("bird")
    except ValueError as exc:
        print(f"factory error: {exc}")


if __name__ == "__main__":
    main()
