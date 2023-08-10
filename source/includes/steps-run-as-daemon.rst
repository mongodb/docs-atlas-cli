.. procedure::
   :style: normal

   .. step:: Start the daemon.

      Run the following command to start the daemon:

      .. code-block::

         docker run -d --name mongodb/atlas mongodb/atlas

   .. step:: Get a shell.

      Run the following command to get a shell:
      
      .. code-block:: 
          
         docker exec -it mongodb/atlas bash

   .. step:: Authenticate and run {+atlas-cli+} commands.

      To authenticate and run commands, follow the steps to
      :ref:`connect from the {+atlas-cli+} <connect-atlas-cli>`, adding
      ``docker exec -it mongodb/atlas`` before each {+atlas-cli+}
      command. For example, to run the :ref:`atlas auth login
      <atlas-auth-login>` command, run the following command:
      
      .. code-block::
         
         docker exec -it mongodb/atlas atlas auth login

      After you authenticate, you can run {+atlas-cli+} commands by
      adding adding ``docker exec -it mongodb/atlas`` before each 
      {+atlas-cli+} command. For example, to run the :ref:`atlas --help <atlas>`
      command, run the following command:

      .. code-block::

         docker exec -it mongodb/atlas atlas --help
