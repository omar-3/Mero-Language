import tkinter as tk
from tkinter import ttk
from tkinter import filedialog
import subprocess
import tempfile
import os


class Root(tk.Tk):
    def __init__(self):
        super(Root, self).__init__()
        self.title("Mero Compiler")
        self.minsize(640, 400)
        self.src_code_file = None
        self.populateUI()

    def populateUI(self):
        self.labelFrame = ttk.LabelFrame(self, text="Options")
        self.labelFrame.grid(column=1, row=3, padx=5, pady=5)
        self.textFrame = ttk.LabelFrame(self, text="Output")
        self.textFrame.grid(column=2, row=3, padx=5, pady=5)

        self.build_button = ttk.Button(
            self.labelFrame, text="Show source code", command=self.showSrcCode
        )
        self.build_button.grid(column=1, row=1, padx=5, pady=5)
        self.browse_button = ttk.Button(
            self.labelFrame, text="Browse", command=self.fileDialog
        )
        self.browse_button.grid(column=1, row=2, padx=5, pady=5)
        self.compile_button = ttk.Button(
            self.labelFrame, text="Compile", command=self.compileSrc
        )
        self.compile_button.grid(column=1, row=3, padx=5, pady=5)

        self.text_area = tk.Text(self.textFrame)
        self.text_area.grid(column=1, row=1)
        self.text_area.configure(state="disabled")

    def fileDialog(self):
        self.src_code_file = filedialog.askopenfilename(
            initialdir=os.getcwd(),
            title="Choose source code file",
            filetypes=(("mero lang source files", "*.mero"), ("all files", "*.*")),
        )
        tk.messagebox.showinfo(message=f"Loaded file: {self.src_code_file}")

    def displayText(self, text):
        self.text_area.configure(state="normal")
        self.text_area.delete('1.0', tk.END)                     # delete current text
        self.text_area.insert('1.0', text)                    # add at line 1, col 0
        self.text_area.mark_set(tk.INSERT, '1.0')                # set insert cursor
        self.text_area.configure(state="disabled")

    def compileSrc(self):
        if not self.src_code_file:
            tk.messagebox.showerror(message="No loaded file")
            return
        self.quadruples_file = tempfile.mktemp(prefix="mero_lang")
        try:
            subprocess.run(
                f"go run ./main.go {self.src_code_file} > {self.quadruples_file}",
                check=True,
                shell=True
            )
        except Exception as e:
            tk.messagebox.showerror(
                message=f"Error: {e}. Make sure you have built the compiler"
            )
            return

        with open(self.quadruples_file) as fd:
            text = fd.read()
            self.displayText(text)
        tk.messagebox.showinfo(message="Compilation complete!")

    def showSrcCode(self):
        if not self.src_code_file:
            tk.messagebox.showerror(message="No loaded file")
            return
        tk.messagebox.showinfo(message="Displaying source code!")
        with open(self.src_code_file) as fd:
            text = fd.read()
            self.displayText(text)



if __name__ == "__main__":
    os.chdir(os.path.dirname(os.path.abspath(__file__)))
    root = Root()
    root.mainloop()
