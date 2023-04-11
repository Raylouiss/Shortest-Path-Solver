from tkinter import *
import tkintermapview
import socket
from tkinter import filedialog
import os
import subprocess

# root = Tk()

# my_label = LabelFrame(root)

# my_label.pack(pady= 20)

# map_widget = tkintermapview.TkinterMapView(my_label, width = 800, height = 600, corner_radius=0)
# # Set Coordinate
# # map_widget.set_position(36.1699, -115.1396)
# # Set a Zoom Level
# map_widget.set_zoom(15)

# map_widget.set_address("Rumah Sakit Umum Imelda Pekerja Indonesia")

# map_widget.pack()
# # input_label = tk.Label(root, text="Input:")
# # input_label.pack()

# # input_entry = tk.Entry(root)
# # input_entry.pack()

# # send_button = tk.Button(root, text="Send", command=send_input_to_go_algorithm)
# # send_button.pack()

# # result_label = tk.Label(root, text="")
# # result_label.pack()

# root.mainloop()
class Application(Frame):
    def __init__(self, master=None):
        # Get the width and height of the screen
        screen_width = root.winfo_screenwidth()
        screen_height = root.winfo_screenheight()

        # Set the size and position of the window
        window_width = 1150
        window_height =648
        x = (screen_width - window_width) // 2
        y = (screen_height - window_height) // 2
        master.geometry(f"{window_width}x{window_height}+{x}+{y}")
        master.configure(bg= '#282B34')
        master.resizable(False, False)
        super().__init__(master)
        self.pack()
        self.create_widgets()

    def send_input_to_go_algorithm(self):
        input_str = self.selected_file_name + "@" + self.radio_var.get() + "@" + self.entry_start.get() + "@" + self.entry_goal.get()
        # Connect to Go Algorithm
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        s.connect(('localhost', 8080))
        s.sendall(input_str.encode())

        # Receive result from Go Algorithm
        result = s.recv(1024).decode()
        result_list = result.split(" ")
        n_result_list= len(result_list)
        result_path = ' '.join(result_list[:n_result_list-1])
        result_cost = result_list[n_result_list-1]
        self.label_result_container_path.config(text="Path:" + result_path)
        self.label_result_container_cost.config(text="Cost:" + result_cost)
        #create path disini
        # self.result_label.config(text=result)
    
    def choose_file_name(self):
        # initialdir = os.path.abspath(os.path.join(os.getcwd(), "..", "main.go"))
        # self.fileName = filedialog.askopenfilenames(initialdir= initialdir, title = "Select A File", filetypes=(("txt files", "*.txt"),("all files", "*.*")))
        # return self.fileName
        subprocess.Popen(['go', 'run', 'main.go']) # Run the Go file
        self.selected_file_name = filedialog.askopenfilename(initialdir="./test", title="Select A File", filetypes=(("txt files", "*.txt"),("all files", "*.*")))
        # self.my_file_label.config(text=self.selected_file_name)
        # return self.selected_file_name
        self.read_file()
        #create mark disini
        # print(self.file_data)
        # n = self.file_data[0]
        # n = int(n)
        # print(n)
        # my_dict = {}
        # print(self.file_data[1])
        # for coordinate in self.file_data[1:n+1]:
        #     print(coordinate)
        #     splitData = coordinate.strip().split(" ")
        #     print(len(splitData))
        #     my_dict[splitData[0]] = (splitData[1], splitData[2])
        # adjMatrix = {}
        # for adj in self.file_data[n+1 :]:
        #     splitAdjData = adj.strip().split("\t")
        #     adjMatrix.append(splitAdjData)
        # for markerKey in my_dict:
        #     latitude, longitude = my_dict[markerKey]
        #     name = markerKey
        #     self.map_widget.set_marker(latitude, longitude, title = name)

    def read_file(self):
        self.my_dic = {}
        self.adjMatrix = {}
        with open(self.selected_file_name, 'r') as file:
            n = int(file.readline())
            for line in file.readlines()[:n]:
                parts = line.split()
                self.my_dic[parts[0]] = (parts[1], parts[2])
            for adj in file.readlines()[n+1 :]:
                splitAdjData = adj.strip().split("\t")
                self.adjMatrix.append(splitAdjData)
            # self.adjMatrix = [line.strip().split('\t') for line in file.readlines()]

        print(self.my_dic)
        print(self.adjMatrix)
                


    # def mark_location(self):

    def create_widgets(self):
        # Create container
        # self.container = Frame(self, bg='gray', width=800, height=400)
        # self.container.pack(side="top", fill="both", expand=True)

        self.left_container = Frame(self, bg="black", width=450, height=700)

        # Choose File Part
        self.file_container = Frame(self.left_container, width=300, height=200, bg= 'black')
        self.label_file_container = Label(self.file_container, text="Choose File", font=("Arial", 20), fg='white', bg='black')
        self.label_file_container.pack(pady=(60, 20))
        self.choose_file_btn = Button(self.file_container)
        self.choose_file_btn["text"] = "Upload File"
        self.choose_file_btn["command"] = self.choose_file_name
        self.choose_file_btn["font"] = ("Courier New", 12)
        self.choose_file_btn["bg"] = 'gray'
        self.choose_file_btn.pack(pady=(10, 30))

        # Choose Input Part
        self.input_container = Frame(self.left_container, width=300, height=150, bg= 'black')
        self.label_input_start = Label(self.input_container, text="Input Starting Point : ", anchor="nw", justify="left", fg= "white", bg= 'black', font=("Courier New", 12))
        self.label_input_start.grid(row=0, column=0, padx=10)
        self.entry_start = Entry(self.input_container)
        self.entry_start.grid(row=0, column=1, padx=10, pady=10)
        self.label_input_goal = Label(self.input_container, text="Input Goal Point     : ", anchor="nw", justify="left", fg= "white", bg= 'black', font=("Courier New", 12))
        self.label_input_goal.grid(row=1, column=0, padx=10, pady=(10, 30))
        self.entry_goal = Entry(self.input_container)
        self.entry_goal.grid(row=1, column=1, padx=10, pady=(10, 30))

        # Choosing Algorithm Part
        self.algorithm_container = Frame(self.left_container, width=300, height=150, bg= 'black')

        # Create radio buttons
        self.radio_var = StringVar()
        self.radio_var.set(None)

        self.radio_button_1 = Radiobutton(self.algorithm_container, text=" A* Algorithm", variable=self.radio_var, value="Option 1", fg= "white", bg= 'black', font=("Courier New", 12), selectcolor="gray")
        self.radio_button_2 = Radiobutton(self.algorithm_container, text="UCS Algorithm", variable=self.radio_var, value="Option 2", fg= "white", bg= 'black', font=("Courier New", 12), selectcolor="gray")

        # Place radio buttons in the frame
        self.radio_button_1.pack()
        self.radio_button_2.pack()

        # Pack child frames of the left container
        self.file_container.pack(side="top", padx=10, pady=10)
        self.input_container.pack(side="top", padx=10, pady=10)
        self.algorithm_container.pack(side="top", padx=10, pady=10)

        # Button container
        self.button_container = Frame(self.left_container, width=300, height=150, bg= "black")
        self.submit_button = Button(self.button_container)
        self.submit_button["text"] = "Search"
        self.submit_button["bg"] = "gray"
        self.submit_button["font"] = ("Courier New", 11)
        self.submit_button["command"] = self.send_input_to_go_algorithm
        self.submit_button.pack()
        self.button_container.pack(padx=10, pady=(20, 10))

        self.left_container.pack(side="left", fill=BOTH)

        # Right Container
        self.right_container = Frame(self, width=850, height=700)

        # Map Container
        self.map_container = Frame(self.right_container, width= 850, height= 600)
        self.map_widget = tkintermapview.TkinterMapView(self.map_container, width = 850, height = 600, corner_radius=0)
        # self.map_widget.set_address("ITB", marker=True)
        # self.map_widget.set_address("Kebun Binatang Bandung", marker=True)
        # self.map_widget.create_line("ITB", "Kebun Binatang Bandung", color="red", width=2)
        self.map_widget.pack()
        self.map_container.pack(side="top")

        # Result Container
        self.result_container = Frame(self.right_container, width= 850, height= 100)
        self.label_result_container_path = Label(self.result_container, text="Path : ", anchor="nw", justify="left", fg= "white", bg= 'gray', font=("Courier New", 12), width= 850)
        self.label_result_container_path.pack()
        self.label_result_container_cost = Label(self.result_container, text="Cost : ", anchor="nw", justify="left", fg= "white", bg= 'gray', font=("Courier New", 12), width= 850)
        self.label_result_container_cost.pack()
        self.result_container.pack()
        self.right_container.pack(side="right", fill=BOTH)

        # self.result_label = Label(self.result_container, text="hasil_path", anchor="nw", justify="left", fg= "white", bg= 'gray', font=("Courier New", 12), width= 850)
        # self.result_label.pack()


        # self.label.place(relx=0.5, rely=0.5, anchor="center")
        
        # self.choose_file_btn = Button(self.left_container)
        # self.choose_file_btn["text"] = "Choose File Btn"
        # self.choose_file_btn["command"] = self.choose_file_name
        # self.choose_file_btn.pack()

        # self.my_file_label = Label(self.left_container, text="")
        # self.my_file_label.pack()

        # self.hi_there = Button(self)
        # self.hi_there["text"] = "Hello World\n(click me)"
        # self.hi_there["command"] = self.send_input_to_go_algorithm
        # self.hi_there.pack(side="top")

        # self.quit = Button(self, text="QUIT", fg="red", command=root.destroy)
        # self.quit.pack(side="bottom")

        # self.result_label = Label(self, text="")
        # self.result_label.pack()
    
    


    # def say_hi(self):
    #     print("hi there, everyone!")

root = Tk()
app = Application(master=root)
app.mainloop()