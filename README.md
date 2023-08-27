
<img align="right" alt="Logo" width="180" height="180" src="https://avatars.githubusercontent.com/u/98979517?s=200&v=4" />

# ReCT

## Important notice
***As of February 2nd 2023 ReCT is no longer maintained and is considered dead.*** The reasons for this are mostly personal as I do not think that I can develop this language into what it deserves to be.  
My vision of "the perfect programming language" has changed drastically over these last few years, continuing development would mean either reworking and removing a lot of language features which make up a core part of the language (resulting in ReCT not being ReCT anymore) or keeping all the older stuff and adding new, totally different feeling things which would make using the language a confusing mess.  

Therefore I have decided to move on to new projects and leave ReCT behind. I want to thank everyone who supported me and this project and helped along the way, no matter how large the contribution, thank you all.

I'll be seeing you around,  
David / Redcube

---

<p>
	  ReCT is a statically typed, easy to use language with a package manager.
    <br />
    <a href="https://docs.rect.ml/"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://rect.ml">Website</a>
    ·
    <a href="https://github.com/ReCT-Lang/ReCT-Go-Compiler/issues">Report Bug</a>
    ·
    <a href="https://github.com/ReCT-Lang/ReCT-Go-Compiler/pulls">Request Feature</a>
</p>

<hr>

<!-- ABOUT THE PROJECT -->
## About The Project

[![Go](https://github.com/ReCT-Lang/ReCT-Go-Compiler/actions/workflows/go.yml/badge.svg)](https://github.com/ReCT-Lang/ReCT-Go-Compiler/actions/workflows/go.yml)  

The first ReCT compiler was written by [RedCubeDev](https://github.com/RedCubeDev-ByteSpace) in C# ([link](https://github.com/RedCubeDev-ByteSpace/ReCT)), this compiler aims to be the successor to
the original ReCT compiler. This compiler is still work in process, many of the features in the rest of this README and on the documentation detail language features in the current version of ReCT (2.3),
this compiler does not support ReCT v2.3 at this current time.

ReCT is a compiled programming language written in Go. It is statically typed, and it's syntax is fast to memorize. 
It has built-in functions that allow users to have a great experience while programming.

Features:
* Optional object-oriented programming
* It is memory safe
* Built-in functions for better productivity
* Statically typed
* Faster than other programming languages

It has a lot more features that can be explored! You can check them out in the [documentation](https://docs.rect.ml/).

## Installation

If you want to try ReCT, the first thing you should do is to install it locally.

1. Clone the repository using `git clone git@github.com:ReCT-Lang/ReCT-Go-Compiler.git` or `git clone https://github.com/ReCT-Lang/ReCT-Go-Compiler.git`
2. Go into the project directory
3. Run `go build -v -a -o "rgoc"` . (requires go to be installed btw)
4. An executable file called `rgoc` will be produced. Move this into an installation directory or anywhere really...
5. If on Windows add as an environmental variable, if on Linux add the executable to your PATH variable.
6. Run the help menu via `rgoc -h`

## Examples

Here is an example of a hello world program in ReCT.
```js
package sys; 

function HelloWorld() string
{
    return "Hello, World";
}

sys::Print(HelloWorld());
```

For more examples, please refer to the [Documentation](https://docs.rect.ml/).


<!-- ROADMAP -->
## Roadmap

- [x] Lexer
- [x] Parser
- [x] Binder
- [x] Lowerer
- [x] Evaluator
- [x] IR emitter

### Version Roadmap
- [x] ReCT v1.1
- [x] ReCT v1.2
- [ ] ReCT v2.0 (skipping all 1.3 Versions)
- [ ] ReCT v2.1
- [ ] ReCT v2.3

See the [open issues](https://github.com/ReCT-Lang/ReCT-Go-Compiler/issues) for a full list of proposed features (and known issues).

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repository and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the GNU.3 License. See `LICENSE` for more information.

## Links

Follow these links and tell us your thoughts, ask questions, and be part of the Rect community!

* [Discord](https://discord.gg/kk9MsnABdF)
* [Website](http://rect.ml/)

## Acknowledgments

Thanks to these special people that made this project happen.

<a href="https://github.com/ReCT-Lang/ReCT-Go-Compiler/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=ReCT-Lang/ReCT-Go-Compiler" />
</a>
